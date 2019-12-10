package faq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDriver() Driver {
	rootQuestion := Question{
		text:   "mulai",
		answer: "apa yang ingin anda ketahui?",
		choices: []Question{
			Question{
				text:   "apa itu e-ktp?",
				answer: "e-ktp adalah sesuatu",
			},
			Question{
				text:   "urus e-tkp",
				answer: "dari mana anda?",
				choices: []Question{
					Question{
						text:   "palu",
						answer: "sudah beli blanko?",
						choices: []Question{
							Question{text: "sudah", answer: "silahkan ke X"},
							Question{text: "belum", answer: "silahkan ke Y"},
						},
					},
					Question{
						text:   "luar palu",
						answer: "minta surat dulu",
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

	assert.Equal(t, "mulai", q.text)
	assert.Equal(t, "apa yang ingin anda ketahui?", r.text)

	r, _ = driver.Ask("apa itu e-ktp?").(Reply)
	assert.Equal(t, "e-ktp adalah sesuatu", r.text)

	r, _ = driver.Ask("sebelumnya").(Reply)
	assert.Equal(t, "apa yang ingin anda ketahui?", r.text)

	r, _ = driver.Ask("urus e-tkp").(Reply)
	assert.Equal(t, "dari mana anda?", r.text)

	r, _ = driver.Ask("luar palu").(Reply)
	assert.Equal(t, "minta surat dulu", r.text)

	r, _ = driver.Ask("ulang").(Reply)
	assert.Equal(t, "minta surat dulu", r.text)

	r, _ = driver.Ask("kembali ke awal").(Reply)
	assert.Equal(t, "apa yang ingin anda ketahui?", r.text)
}
