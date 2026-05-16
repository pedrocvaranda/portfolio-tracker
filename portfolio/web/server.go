package web

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
	"github.com/pedrocvaranda/portfolio-tracker/portfolio/montecarlo"
)

//go:embed assets/*
var assetsFS embed.FS

// Server exposes the portfolio tracker as a web application.
type Server struct {
	engine montecarlo.Engine
	index  *template.Template
}

type pageData struct {
	GeneratedAt string
}

type simulateRequest struct {
	Ticker     string  `json:"ticker"`
	StartPrice float64 `json:"startPrice"`
	Drift      float64 `json:"drift"`
	Volatility float64 `json:"volatility"`
	Days       int     `json:"days"`
	Paths      int     `json:"paths"`
	Seed       int64   `json:"seed"`
}

type simulateResponse struct {
	Ticker       string      `json:"ticker"`
	StartPrice   float64     `json:"startPrice"`
	ExpectedEnd  float64     `json:"expectedEnd"`
	Percentile05 float64     `json:"percentile05"`
	Percentile50 float64     `json:"percentile50"`
	Percentile95 float64     `json:"percentile95"`
	Paths        [][]float64 `json:"paths"`
	Days         int         `json:"days"`
	Count        int         `json:"count"`
	GeneratedAt  string      `json:"generatedAt"`
}

func NewServer() (*Server, error) {
	index, err := template.ParseFS(assetsFS, "assets/index.html")
	if err != nil {
		return nil, err
	}
	return &Server{engine: montecarlo.NewEngine(), index: index}, nil
}

func MustNewServer() *Server {
	server, err := NewServer()
	if err != nil {
		panic(err)
	}
	return server
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/api/simulate", s.handleSimulate)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS()))))
	return mux
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = s.index.Execute(w, pageData{GeneratedAt: time.Now().Format("2006-01-02 15:04")})
}

func (s *Server) handleSimulate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req simulateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}
	params := montecarlo.Params{
		Ticker:     strings.ToUpper(strings.TrimSpace(req.Ticker)),
		StartPrice: req.StartPrice,
		Drift:      req.Drift,
		Volatility: req.Volatility,
		Days:       req.Days,
		Paths:      req.Paths,
		Seed:       req.Seed,
	}
	result, err := s.engine.Run(params)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	resp := simulateResponse{
		Ticker:       result.Ticker,
		StartPrice:   result.StartPrice,
		ExpectedEnd:  result.ExpectedEnd,
		Percentile05: result.Percentile05,
		Percentile50: result.Percentile50,
		Percentile95: result.Percentile95,
		Paths:        samplePaths(result.Paths, 18),
		Days:         params.WithDefaults().Days,
		Count:        result.Count(),
		GeneratedAt:  result.GeneratedAt.Format(time.RFC3339),
	}
	writeJSON(w, http.StatusOK, resp)
}

func staticFS() fs.FS {
	static, err := fs.Sub(assetsFS, "assets/static")
	if err != nil {
		panic(err)
	}
	return static
}

func samplePaths(paths []models.SimulationPath, limit int) [][]float64 {
	if limit > len(paths) {
		limit = len(paths)
	}
	sampled := make([][]float64, limit)
	for i := 0; i < limit; i++ {
		sampled[i] = append([]float64(nil), paths[i].Prices...)
	}
	return sampled
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
