# Email Verifier Tool

This email verifier tool will check  MX (Mail Exchange) and security checks such as SPF (Sender Policy Framework) and DMARC (Domain-based Message Authentiation, Reporting and Conformance) while verifying the given mail server.
## Installation


```bash
go get github.com/peterhyperaung/emailverifier
```
    
## License
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)


## Documentation

#### Import the package
```go
import "github.com/peterhyperaung/emailverifier"
```

#### Check Email

```go
// CheckEmail(email string) bool
emailverifier.CheckEmail("example@gmail.com")
```

#### Validate Email Format

```go
// ValidateEmailFormat(email string) bool
emailverifier.ValidateEmailFormat("example@gmail.com")
```



