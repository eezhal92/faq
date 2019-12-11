package faq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDriver() *Driver {
	rootQuestion := Question{
		Text:   "mulai",
		Answer: "apa yang ingin anda ketahui?",
		Choices: []Question{
			Question{
				Text:   "apa itu e-ktp?",
				Answer: "e-ktp adalah sesuatu",
			},
			Question{
				Text:   "urus e-tkp",
				Answer: "dari mana anda?",
				Choices: []Question{
					Question{
						Text:   "palu",
						Answer: "sudah beli blanko?",
						Choices: []Question{
							Question{Text: "sudah", Answer: "silahkan ke X"},
							Question{Text: "belum", Answer: "silahkan ke Y"},
						},
					},
					Question{
						Text:   "luar palu",
						Answer: "minta surat dulu",
					},
				},
			},
		},
	}

	return NewDriver(rootQuestion, "sebelumnya", "ulang", "kembali ke awal")
}

func TestFullConvoFlow(t *testing.T) {
	driver := createDriver()

	r := driver.Boot()
	q, _ := driver.CurrentQuestion().(Question)

	assert.Equal(t, "mulai", q.Text)
	assert.Equal(t, "apa yang ingin anda ketahui?", r.Text)

	r = driver.Ask("apa itu e-ktp?")
	assert.Equal(t, "e-ktp adalah sesuatu", r.Text)

	r = driver.Ask("sebelumnya")
	assert.Equal(t, "apa yang ingin anda ketahui?", r.Text)

	r = driver.Ask("urus e-tkp")
	assert.Equal(t, "dari mana anda?", r.Text)

	r = driver.Ask("luar palu")
	assert.Equal(t, "minta surat dulu", r.Text)

	r = driver.Ask("ulang")
	assert.Equal(t, "minta surat dulu", r.Text)

	r = driver.Ask("kembali ke awal")
	assert.Equal(t, "apa yang ingin anda ketahui?", r.Text)
}
