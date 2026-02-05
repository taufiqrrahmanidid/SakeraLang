# SakeraLang 🌴

![SakeraLang](https://img.shields.io/badge/version-0.2.0-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)

**Bahasa pemrograman berbasis bahasa Madura** - Bringing the beauty of Madurese language to programming!

SakeraLang adalah bahasa pemrograman yang menggunakan kata kunci berbahasa Madura, terinspirasi dari bahasa pemrograman lokal Indonesia lainnya seperti PrabogoLang, SundaLang, dan lainnya.

---

## ✨ Fitur

### ✅ Fitur Lengkap (v0.2.0)

- **Variables** - Deklarasi variabel dengan `sango`
- **Data Types** - Integer, String, Boolean
- **Operators** - Aritmatika (+, -, *, /), Perbandingan (<, >, ==, !=)
- **Comments** - Komentar dengan `//`
- **Conditionals** - `mon` (if) dan `laen` (else)
- **Loops** - `selama` (while) dan `ulang` (for)
- **Functions** - `fungsi` dengan parameter dan `mareh` (return)
- **Recursion** - Dukungan penuh untuk rekursi
- **Input/Output** - `toles` (print) dan `macah` (read)

---

## 📚 Kata Kunci (Keywords)

| Keyword | Arti Indonesia | Arti English | Contoh |
|---------|---------------|--------------|--------|
| `sango` | variabel | variable | `sango nama = "Taufiq"` |
| `toles` | cetak | print | `toles nama` |
| `macah` | baca | read/input | `macah nama` |
| `mon` | jika | if | `mon x > 10 { }` |
| `laen` | lainnya | else | `laen { }` |
| `selama` | selama | while | `selama i < 5 { }` |
| `ulang` | ulang | for | `ulang sango i = 0; i < 5; sango i = i + 1 { }` |
| `fungsi` | fungsi | function | `fungsi tambah(a, b) { }` |
| `mareh` | kembali | return | `mareh hasil` |
| `bender` | benar | true | `sango aktif = bender` |
| `sala` | salah | false | `sango nonaktif = sala` |

---

## 🚀 Instalasi

### Prerequisites
- Go 1.21 atau lebih baru

### Build dari Source
```bash
git clone https://github.com/taufiqrrahmanidid/SakeraLang.git
cd SakeraLang
go build -o sakera ./cmd/sakera
```

Atau gunakan Makefile:
```bash
make build
```

---

## 💻 Cara Menggunakan

### Mode Interaktif (REPL)
```bash
./sakera
```

Output:
```
=================================
  Selamet Rabu e  SakeraLang v0.2.0
  Bahasa Pemrograman Madura
=================================
Ketik 'keluar' untuk exit

sakera> sango nama = "Madura"
sakera> toles nama
Madura
sakera> keluar
Mator Sakalangkong!
```

### Jalankan File
```bash
./sakera examples/hello.sakera
```

---

## 📖 Contoh Program

### 1. Hello World
```sakera
sango nama = "Madura"
toles nama
```

### 2. Variabel dan Operasi
```sakera
sango a = 10
sango b = 5
sango hasil = a + b
toles hasil  // Output: 15
```

### 3. Conditional
```sakera
sango umur = 20

mon umur > 17 {
    toles "Dewasa"
}
laen {
    toles "Anak-anak"
}
```

### 4. While Loop
```sakera
sango i = 0

selama i < 5 {
    toles i
    sango i = i + 1
}
```

### 5. For Loop
```sakera
ulang sango i = 0; i < 5; sango i = i + 1 {
    toles i
}
```

### 6. Functions
```sakera
fungsi tambah(a, b) {
    mareh a + b
}

sango hasil = tambah(5, 3)
toles hasil  // Output: 8
```

### 7. Fibonacci (Recursion)
```sakera
fungsi fibonacci(n) {
    mon n < 2 {
        mareh n
    }
    laen {
        sango a = fibonacci(n - 1)
        sango b = fibonacci(n - 2)
        mareh a + b
    }
}

toles fibonacci(10)  // Output: 55
```

### 8. Input dari User
```sakera
toles "Masukkan nama Anda:"
macah nama
toles "Halo, "
toles nama
```

---

## 📂 Struktur Project
```
SakeraLang/
├── cmd/
│   └── sakera/
│       └── main.go          # Entry point program
├── pkg/
│   ├── lexer/
│   │   ├── token.go         # Token definitions
│   │   └── lexer.go         # Lexical analyzer
│   ├── parser/
│   │   ├── ast.go           # Abstract Syntax Tree
│   │   └── parser.go        # Parser
│   └── evaluator/
│       └── evaluator.go     # Interpreter
├── examples/                # Contoh program
│   ├── hello.sakera
│   ├── fibonacci.sakera
│   ├── comprehensive_all.sakera
│   └── ...
├── README.md
├── LICENSE
└── Makefile
```

---

## 🎯 Roadmap

### ✅ v0.2.0 (Current)
- ✅ Comments
- ✅ Conditionals (mon/laen)
- ✅ Loops (selama/ulang)
- ✅ Functions & Recursion
- ✅ Return values
- ✅ Input/Output

### 🔜 v0.3.0 (Future)
- [ ] Arrays/Lists
- [ ] String methods
- [ ] Logical operators (dan/atau)
- [ ] Break & Continue
- [ ] Error handling
- [ ] Standard library
- [ ] File I/O
- [ ] Better error messages

---

## 🤝 Kontribusi

Kontribusi sangat diterima! Silakan:

1. Fork repository ini
2. Buat branch baru (`git checkout -b fitur-baru`)
3. Commit perubahan (`git commit -am 'Tambah fitur baru'`)
4. Push ke branch (`git push origin fitur-baru`)
5. Buat Pull Request

Lihat [CONTRIBUTING.md](CONTRIBUTING.md) untuk panduan lengkap.

---

## 📜 Lisensi

MIT License - Lihat file [LICENSE](LICENSE) untuk detail.

---

## 🙏 Terima Kasih

Terinspirasi dari komunitas bahasa pemrograman lokal Indonesia:
- [PrabogoLang](https://github.com/prabowo/PrabogoLang)
- [SundaLang](https://github.com/sundanese/SundaLang)
- Dan komunitas open source Indonesia lainnya

---

## 📞 Kontak

- **GitHub:** [@taufiqrrahmanidid](https://github.com/taufiqrrahmanidid)
- **Repository:** [SakeraLang](https://github.com/taufiqrrahmanidid/SakeraLang)

---

<div align="center">

**Dibuat dengan ❤️ untuk Madura**

⭐ Star repository ini jika Anda suka SakeraLang!

</div>
