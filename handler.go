package main

import (
	"log"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	geoDBPath, proAddr string

	geoDB          *geoip2.Reader
	requestCounter *prometheus.CounterVec
	metaPool       sync.Pool
}

func NewHandler(geoDBPath string, pro string) (h *Handler, err error) {

	h = &Handler{geoDBPath: geoDBPath, proAddr: pro}

	h.metaPool = sync.Pool{New: func() interface{} {
		return &Meta{}
	}}

	if h.geoDBPath != "" {
		h.geoDB, err = geoip2.Open(h.geoDBPath)
		if err != nil {
			log.Print(err)
		}
		log.Printf("Geo Path %s", h.geoDBPath)
	}

	if h.proAddr != "" {
		h.requestCounter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "http",
				Subsystem: "requests",
				Name:      "total",
				Help:      "The total number of http request",
			}, []string{"type"})

		prometheus.MustRegister(h.requestCounter)
		http.Handle("/metrics", promhttp.Handler())
		go http.ListenAndServe(h.proAddr, nil)
		log.Println("metric export at ", h.proAddr)
	}

	return
}

func (h *Handler) PrometheusHandler(c *gin.Context) {

}

func (h *Handler) Root(c *gin.Context) {

	accept := c.Request.Header.Get("Accept")
	if len(accept) < 4 {
		h.RootRaw(c)
		return
	}

	al := strings.SplitN(accept, ",", 2)

	switch al[0] {
	case binding.MIMEJSON:
		h.requestCounter.WithLabelValues("json").Inc()
		h.RootJson(c)
	case binding.MIMEHTML:
		h.requestCounter.WithLabelValues("html").Inc()
		h.RootHTML(c)
	default:
		h.requestCounter.WithLabelValues("raw").Inc()
		h.RootRaw(c)
	}
}

func (h *Handler) RootHTML(c *gin.Context) {
	m := h.makeMeta(c)
	defer h.metaPool.Put(m)

	switch c.Query("lang") {
	case "cat":
		c.HTML(http.StatusOK, "index-cat.tpl", m)
	default:
		c.HTML(http.StatusOK, "index.tpl", m)
	}
}

func (h *Handler) RootRaw(c *gin.Context) {
	c.String(http.StatusOK, "%s\n", c.ClientIP())
}

func (h *Handler) RootJson(c *gin.Context) {
	m := h.makeMeta(c)
	defer h.metaPool.Put(m)
	c.IndentedJSON(http.StatusOK, m)
}

type Meta struct {
	IP      string `json:"ip"`
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

func (h *Handler) makeMeta(c *gin.Context) (m *Meta) {

	m = h.metaPool.Get().(*Meta)
	m.IP = c.ClientIP()

	if h.geoDB == nil {
		return
	}

	ip := net.ParseIP(m.IP)
	if country, err := h.geoDB.Country(ip); err == nil {
		m.Country = country.Country.Names[acceptCode]
	} else {
		log.Print(err)
	}
	if city, err := h.geoDB.City(ip); err == nil {
		m.City = city.City.Names[acceptCode]
	} else {
		log.Print(err)
	}
	return
}
