package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "%d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// for k, v := range r.Header {
	// fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// }

	// fmt.Fprintf(w, "Host = %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cycles, size, nframes, delay := 5, 100, 64, 8
	res := 0.001
	for k, v := range r.Form {
		if k == "cycles" && len(v) == 1 {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				cycles = 5
			} else {
				cycles = i
			}
		}
		if k == "size" && len(v) == 1 {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				size = 100
			} else {
				size = i
			}
		}
		if k == "nframes" && len(v) == 1 {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				nframes = 64
			} else {
				nframes = i
			}
		}
		if k == "delay" && len(v) == 1 {
			i, err := strconv.Atoi(v[0])
			if err != nil {
				delay = 8
			} else {
				delay = i
			}
		}
		if k == "res" && len(v) == 1 {
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				res = 0.001
			} else {
				res = f
			}
		}
		// fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

var palette = []color.Color{color.Black, color.RGBA{R: 0xFF, A: 0xFF}, color.RGBA{G: 0xFF, A: 0xFF}, color.RGBA{B: 0xFF, A: 0xFF}}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	n := uint8(rand.Intn(len(palette)-1) + 1)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), n)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
