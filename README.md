# 🔍 Header Inspector

**Header Inspector** is a simple Go-based CLI tool that analyzes the HTTP response headers of a given URL — with a focus on **security-related headers**.

---

## 📦 Features

- Accepts a URL and inspects its HTTP response headers
- Checks for presence of common security headers:
  - `Content-Security-Policy`
  - `Strict-Transport-Security`
  - `X-Content-Type-Options`
  - `X-Frame-Options`
  - `Referrer-Policy`
  - `Permissions-Policy`
  - `Access-Control-Allow-Origin`
  - `Set-Cookie` (Secure/HttpOnly check)
- Clean CLI output
- Built-in test case for `https://example.com`
- GitHub Actions CI setup for build & test

---

## 🚀 Getting Started

### 🔧 Requirements
- Go 1.20 or later

### 📥 Clone & Build
```bash
git clone https://github.com/Cyb3r3x3r/header-inspector.git
cd header-inspector
go build -o header-inspector ./cmd
```

### 📥 Sample Usage
```shell
go run ./cmd https://example.com
```

#### Output
```
Security Headers for https://example.com:
  Content-Security-Policy:        ✗ Missing
  Strict-Transport-Security:      ✓ Found
  X-Content-Type-Options:         ✓ Found
  X-Frame-Options:                ✗ Missing
  Referrer-Policy:                ✗ Missing
  Permissions-Policy:             ✗ Missing
  Access-Control-Allow-Origin:    ✓ Found
  Set-Cookie:                     ⚠ Insecure Flags
```