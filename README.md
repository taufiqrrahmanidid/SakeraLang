# SakeraLang 🌴

Bahasa pemrograman berbasis bahasa Madura - Bringing the beauty of Madurese language to programming!

## Tentang SakeraLang

SakeraLang adalah bahasa pemrograman yang menggunakan kata kunci berbahasa Madura, terinspirasi dari bahasa pemrograman lokal Indonesia lainnya seperti PrabogoLang dan SundaLang.

## Kata Kunci (Keywords)

- sango - variabel (variable)

- mon - jika (if)

- laen - lainnya (else)

- selama - selama (while)

- ulang - ulang (for)

- mareh - kembali (return)

- fungsi - fungsi (function)

- bender - benar (true)

- sala - salah (false)

- toles - cetak (print)

- macah - baca input (read)

## Instalasi

### Prerequisites

- Go 1.20 atau lebih baru

### Build dari Source

```bash

git clone https://github.com/taufiqrrahmanidid/SakeraLang.git

cd SakeraLang

go build -o sakera cmd/sakera/main.go

```

## Cara Menggunakan

### Mode Interaktif (REPL)

```bash

./sakera

```

### Contoh Kode

```sakera

sango nama = "Madura"

toles nama

sango angka = 10

sango hasil = angka + 5

toles hasil

```

## Contoh Program

Lihat folder examples/ untuk contoh program lebih lengkap.

## Roadmap

- [x] Lexer dan Tokenizer

- [x] Evaluator dasar

- [ ] Parser lengkap

- [ ] Conditional statements (mon/laen)

- [ ] Loops (selama/ulang)

- [ ] Functions (fungsi)

- [ ] File execution

## Kontribusi

Kontribusi sangat diterima! Silakan buat Pull Request atau buka Issue untuk diskusi.

## Lisensi

MIT License - Lihat file LICENSE untuk detail

## Terima Kasih

Terinspirasi dari komunitas bahasa pemrograman lokal Indonesia.

---

Dibuat dengan ❤️ untuk Madura

