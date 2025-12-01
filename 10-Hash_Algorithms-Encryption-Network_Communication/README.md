# Cryptography & Network Communication: A Comprehensive Guide

A detailed exploration of **Hash Algorithms**, **Symmetric & Asymmetric Encryption**, and **Network Communication Protocols**. This repository serves as an educational resource for understanding the cryptographic foundations of modern secure communication.

---

## Table of Contents

1. [Hash Algorithms](#hash-algorithms)
2. [Symmetric Encryption](#symmetric-encryption)
3. [Asymmetric Encryption](#asymmetric-encryption)
4. [Network Communication](#network-communication)
5. [Practical Applications](#practical-applications)
6. [Security Threats & Prevention](#security-threats--prevention)
7. [Implementation Best Practices](#implementation-best-practices)

---

## Hash Algorithms

### Overview

**Cryptographic hash functions** are fundamental building blocks in modern security systems. A hash function takes variable-length input and produces a fixed-size output (the **hash** or **message digest**). The process is deterministic—the same input always produces the same hash.

### Key Properties

| Property | Description |
|----------|-------------|
| **Deterministic** | Same input always produces the same output |
| **Pre-Image Resistant** | Nearly impossible to reverse-engineer the original input from the hash |
| **Collision-Resistant** | Two different inputs should never produce the same hash |
| **Avalanche Effect** | Small changes in input produce completely different hashes |
| **Computationally Efficient** | Fast to compute but computationally secure |

### Common Hash Algorithms

#### MD5 (Message Digest Algorithm 5)

- **Output Size**: 128 bits (16 bytes)
- **Status**: ⚠️ Cryptographically broken
- **Vulnerabilities**: Collision attacks can be generated in less than a minute on modern hardware
- **Current Use**: File integrity verification (non-critical), checksums
- **Recommendation**: **DO NOT USE** for security-critical applications

**Example:**
```
MD5("password") = 5f4dcc3b5aa765d61d8327deb882cf99
MD5("Password") = dc647eb65e6711e155375218212b3964
```

#### SHA Family (Secure Hash Algorithm)

##### SHA-1 (Deprecated)
- **Output Size**: 160 bits (20 bytes)
- **Status**: ❌ Vulnerable to collision attacks
- **Recommendation**: Deprecated for cryptographic purposes

##### SHA-2 (Recommended)
- **Variants**: SHA-224, SHA-256, SHA-384, SHA-512
- **Output Size**: 224, 256, 384, or 512 bits
- **Security**: ✅ Collision-resistant and one-way
- **Status**: Gold standard for secure applications
- **Applications**: TLS/SSL, SSH, blockchain, digital signatures, password hashing

**Example:**
```
SHA-256("password") = 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8
SHA-256("Password") = e7cf3ef4f17c3999a94f2c6f612e8a888e5b1026878e4e19398b23dd3f5c1c16
```

##### SHA-3 (Latest Standard)
- **Output Size**: Configurable (224, 256, 384, 512 bits)
- **Design**: Based on Keccak algorithm (sponge construction)
- **Security**: Newer alternative with different mathematical approach
- **Use Case**: Emerging standard, used in modern systems

### Use Cases for Hashing

1. **Data Integrity Verification**: Detect file corruption or tampering
2. **Password Storage**: Never store plaintext passwords; store hashes instead
3. **Digital Signatures**: Sign documents using hashed content
4. **Merkle Trees**: Organize and verify blockchain data efficiently
5. **Message Authentication**: Combined with keys to create HMAC
6. **Checksums**: Verify downloaded files haven't been corrupted

### Comparison: MD5 vs SHA-256

| Aspect | MD5 | SHA-256 |
|--------|-----|---------|
| **Hash Length** | 128 bits | 256 bits |
| **Security** | Broken ❌ | Strong ✅ |
| **Speed** | Faster | Slightly slower |
| **Collision Resistance** | Vulnerable | Resistant |
| **Recommended Use** | Legacy only | All security purposes |

---

## Symmetric Encryption

### Overview

**Symmetric encryption** uses a **single shared key** to both encrypt and decrypt data. Both parties must possess the same secret key, kept confidential from unauthorized users.

### Key Characteristics

| Characteristic | Details |
|---|---|
| **Key Size** | Typically 128, 192, or 256 bits |
| **Speed** | Fast (~1000x faster than asymmetric) |
| **Scalability** | Efficient for large data volumes |
| **Drawback** | Key distribution challenge |
| **Use Case** | Bulk data encryption |

### Common Symmetric Algorithms

#### AES (Advanced Encryption Standard)

- **Key Sizes**: 128, 192, 256 bits
- **Block Size**: 128 bits (fixed)
- **Type**: Block cipher
- **Status**: ✅ Current U.S. government standard
- **Hardware Support**: AES-NI acceleration on modern CPUs
- **Performance**: Excellent, especially with hardware acceleration

**AES Modes:**
- **ECB** (Electronic Codebook): Avoid! Deterministic leaks patterns
- **CBC** (Cipher Block Chaining): Requires unique IV for each encryption
- **GCM** (Galois/Counter Mode): ✅ Recommended - provides authenticated encryption
- **CTR** (Counter Mode): Stream-like mode, parallelizable

**Example (OpenSSL):**
```bash
# Encrypt with AES-256-GCM
openssl enc -aes-256-gcm -in plaintext.txt -out ciphertext.bin -k password

# Decrypt
openssl enc -aes-256-gcm -d -in ciphertext.bin -out plaintext.txt -k password
```

#### DES & 3DES (Deprecated)

- **DES**: 56-bit key (too short, vulnerable to brute force)
- **3DES**: Applies DES three times with different keys
- **Status**: ❌ Deprecated; replaced by AES
- **Speed**: Slower than AES

#### ChaCha20 / XChaCha20 (Modern Alternative)

- **Key Size**: 256 bits
- **Type**: Stream cipher (encrypts data sequentially)
- **Nonce Size**: ChaCha20 (96 bits), XChaCha20 (192 bits)
- **Performance**: Excellent software implementation (no hardware dependency)
- **Advantage**: Better for embedded systems and mobile devices without AES-NI
- **Status**: ✅ Growing adoption, especially in mobile and privacy-focused apps

**Example (libsodium):**
```c
crypto_stream_chacha20_xor(output, input, nonce, key);
```

#### Comparison: AES vs ChaCha20

| Aspect | AES-256 | ChaCha20-256 |
|--------|---------|--------------|
| **Type** | Block cipher | Stream cipher |
| **Hardware Dependent** | Yes (AES-NI) | No |
| **Performance (with HW)** | Faster | Slower |
| **Performance (without HW)** | Slower | Faster |
| **Mobile Friendly** | Variable | Better |
| **Nonce Reuse** | Dangerous (CBC/GCM) | Dangerous |
| **Adoption** | Ubiquitous ✅ | Growing |

---

## Asymmetric Encryption

### Overview

**Asymmetric encryption**, also called **public-key cryptography**, uses a pair of mathematically linked keys:
- **Public Key**: Shared with everyone; encrypts data
- **Private Key**: Kept secret by owner; decrypts data

### Key Characteristics

| Characteristic | Details |
|---|---|
| **Key Pair** | Public and private keys are mathematically linked |
| **Key Size** | 2048-4096 bits (RSA) or smaller (ECC) |
| **Speed** | Slow (~1000x slower than symmetric) |
| **Key Distribution** | No secure channel needed to share public key |
| **Use Case** | Key exchange, digital signatures, authentication |

### How Asymmetric Encryption Works

**Basic Process:**
1. Alice generates a key pair (public key `pub_A`, private key `priv_A`)
2. Alice publishes her public key
3. Bob encrypts a message using Alice's public key: `ciphertext = encrypt(message, pub_A)`
4. Alice decrypts with her private key: `message = decrypt(ciphertext, priv_A)`

### Common Asymmetric Algorithms

#### RSA (Rivest-Shamir-Adleman)

- **Security Basis**: Difficulty of factoring large prime numbers
- **Key Size**: 2048 bits (minimum recommended), 4096 bits (strong)
- **Applications**: Encryption, digital signatures, key exchange
- **Status**: ✅ Industry standard, widely used
- **Drawback**: Relatively slow, requires large key sizes

**RSA Key Generation:**
```
1. Select two large prime numbers (p, q)
2. Compute n = p × q
3. Compute φ(n) = (p-1) × (q-1)
4. Select e (usually 65537) where gcd(e, φ(n)) = 1
5. Compute d: e × d ≡ 1 (mod φ(n))
6. Public Key = (e, n), Private Key = (d, n)
```

#### ECC (Elliptic Curve Cryptography)

- **Security Basis**: Elliptic Curve Discrete Logarithm Problem (ECDLP)
- **Key Advantage**: Equivalent security with smaller keys (256-bit ECC ≈ 2048-bit RSA)
- **Performance**: Faster than RSA with equivalent security
- **Variants**: ECDSA (digital signatures), ECDH (key exchange)
- **Status**: ✅ Modern standard, increasingly adopted
- **Use Cases**: TLS 1.3, Bitcoin, modern cryptographic protocols

**Comparison: RSA vs ECC**

| Aspect | RSA | ECC |
|--------|-----|-----|
| **256-bit Equivalent** | 2048-bit RSA | 256-bit ECC |
| **Speed** | Slower | Faster |
| **Key Size** | Large | Small |
| **Implementation** | Well-established | Emerging |
| **Patent Issues** | Expired | Some variants |
| **Adoption** | Ubiquitous | Growing |

#### Diffie-Hellman Key Exchange

- **Purpose**: Securely exchange a shared secret over an insecure channel
- **Basis**: Discrete logarithm problem
- **Output**: Shared secret (session key) for symmetric encryption
- **Variants**: DHKE, ECDH (elliptic curve)

**DHKE Process:**
```
1. Alice & Bob agree on prime (p) and generator (g)
2. Alice generates secret (a), sends g^a mod p to Bob
3. Bob generates secret (b), sends g^b mod p to Alice
4. Alice computes: shared_secret = (g^b)^a mod p = g^(ab) mod p
5. Bob computes: shared_secret = (g^a)^b mod p = g^(ab) mod p
6. Both have the same shared secret without revealing a or b
```

### Hybrid Cryptosystems

Modern protocols combine **asymmetric and symmetric** encryption:

1. **Asymmetric** (RSA/ECC): Secure exchange of symmetric key
2. **Symmetric** (AES/ChaCha20): Fast encryption of actual data

**Example: TLS/SSL Handshake**
```
1. Client & Server use RSA/ECDH to exchange symmetric session key
2. Subsequent communication encrypted with AES-GCM (symmetric)
3. Asymmetric overhead only for initial handshake
4. Benefit: Security of asymmetric + speed of symmetric
```

---

## Network Communication

### OSI Model Layers

| Layer | Name | Protocol Examples | Function |
|-------|------|---|---|
| 7 | Application | HTTP, HTTPS, FTP, SSH, DNS, SMTP | User applications, APIs |
| 6 | Presentation | SSL/TLS, encryption, compression | Data formatting, encryption |
| 5 | Session | SSH, TLS, RPC | Connection management |
| 4 | Transport | TCP, UDP, TLS | End-to-end delivery |
| 3 | Network | IP (IPv4/IPv6), ICMP | Routing, logical addressing |
| 2 | Data Link | Ethernet, PPP, ARP | MAC addressing, framing |
| 1 | Physical | Fiber, Copper, Radio | Bits, signals, hardware |

### Transport Layer Protocols

#### TCP (Transmission Control Protocol)

- **Type**: Connection-oriented, reliable
- **Guarantees**: 
  - Data arrives in correct order
  - No packet loss
  - Error detection and correction
- **Overhead**: Higher due to handshaking and acknowledgment
- **Use Cases**: Web browsing (HTTP/HTTPS), Email (SMTP, POP3), File transfer (FTP)
- **Handshake**: Three-way handshake (SYN, SYN-ACK, ACK)

**TCP Connection Process:**
```
Client                          Server
  |-------- SYN (seq=x) -------->|
  |<----- SYN-ACK (seq=y, ack=x+1)---|
  |-------- ACK (seq=x+1, ack=y+1) ->|
  |<====== Connection Established ====|
```

#### UDP (User Datagram Protocol)

- **Type**: Connectionless, unreliable
- **Characteristics**:
  - No connection setup (lower latency)
  - No guaranteed delivery
  - No ordering guarantee
  - Lower overhead
- **Performance**: ~100x faster than TCP
- **Trade-off**: Speed over reliability
- **Use Cases**: Real-time applications, live streaming, online gaming, DNS queries

**UDP vs TCP:**

| Feature | TCP | UDP |
|---------|-----|-----|
| **Connection** | Established (3-way) | Connectionless |
| **Reliability** | Guaranteed delivery | Best effort |
| **Ordering** | Guaranteed order | No guarantee |
| **Speed** | Slower | Faster ⚡ |
| **Overhead** | Higher | Lower |
| **Use Case** | Accuracy critical | Speed critical |

### Network Layer: IP Protocols

#### IPv4 (Internet Protocol version 4)

- **Address Space**: 32 bits (4 octets)
- **Capacity**: ~4.3 billion unique addresses
- **Format**: `192.168.1.1`
- **Header Size**: 20 bytes (minimum)
- **NAT Support**: Necessary due to address exhaustion
- **Status**: Still widely used

#### IPv6 (Internet Protocol version 6)

- **Address Space**: 128 bits (16 octets)
- **Capacity**: 2^128 unique addresses (virtually unlimited)
- **Format**: `2001:0db8:85a3::8a2e:0370:7334`
- **Benefits**:
  - Simpler header processing
  - Built-in IPSec support
  - Better routing efficiency
- **Adoption**: Gradually increasing
- **Status**: Future-proof standard

### Application Layer Protocols

#### HTTP vs HTTPS

| Aspect | HTTP | HTTPS |
|--------|------|-------|
| **Encryption** | None ❌ | TLS/SSL ✅ |
| **Data Protection** | Unencrypted | Encrypted |
| **Authentication** | None | Certificate-based |
| **Port** | 80 | 443 |
| **Man-in-the-Middle** | Vulnerable | Protected |
| **Current Use** | Deprecated | Standard ✅ |

#### DNS (Domain Name System)

- **Purpose**: Translate domain names to IP addresses
- **Port**: UDP 53 (standard queries), TCP 53 (zone transfers)
- **Query Process**:
  1. User enters `google.com` in browser
  2. Browser queries DNS resolver on port 53 (UDP)
  3. DNS resolver queries root nameserver
  4. Returns authoritative nameserver address
  5. Returns IP address to client
  6. Browser connects to IP address

**DNS Security Threats:**
- DNS Spoofing: Attacker returns false IP
- DNS Tunneling: Hiding data in DNS queries
- DNS Cache Poisoning: Injecting false records
- DDoS attacks on port 53

**DNS Security Solutions:**
- DNSSEC: Digital signatures for DNS
- DNS over HTTPS (DoH): Encrypts queries over HTTPS
- DNS over TLS: Port 853 with TLS encryption

#### SSH (Secure Shell)

- **Purpose**: Secure remote access, file transfer, tunneling
- **Port**: 22 (TCP)
- **Authentication**: Public-key cryptography, password, or both
- **Encryption**: Hybrid (asymmetric for key exchange, symmetric for data)
- **Key Exchange**: Diffie-Hellman or ECDH
- **Algorithms**: RSA, ECDSA, Ed25519 for keys; AES, ChaCha20 for encryption

---

## Practical Applications

### TLS/SSL Handshake (HTTPS)

**Complete TLS 1.2 Handshake:**

```
Client                                          Server
  |                                               |
  |------ ClientHello ---------------------->|
  | (supported ciphers, TLS version, random)      |
  |                                               |
  |<------ ServerHello, Certificate -------|
  | (chosen cipher, TLS version, random,          |
  |  server certificate with public key)          |
  |                                               |
  |------ ClientKeyExchange ---------------->|
  | (premaster secret encrypted with pub key)    |
  |                                               |
  |------ ChangeCipherSpec, Finished ------>|
  | (signals use of symmetric encryption)        |
  |                                               |
  |<------ ChangeCipherSpec, Finished ------|
  |                                               |
  |<===== Encrypted Communication =====>|
  | (AES-256-GCM or ChaCha20-Poly1305)          |
  |                                               |
```

**Security Steps:**
1. **Authentication**: Server proves identity via certificate
2. **Key Exchange**: Client generates session key securely
3. **Symmetric Encryption**: Both use session key for fast encryption
4. **Integrity**: HMAC ensures data hasn't been modified

### Digital Signatures

**Process:**

```
Signer (Alice):
  1. Create message
  2. Hash message: h = SHA-256(message)
  3. Encrypt hash with private key: sig = encrypt(h, private_key)
  4. Send message + signature

Verifier (Bob):
  1. Receive message + signature
  2. Decrypt signature with Alice's public key: h' = decrypt(sig, public_key)
  3. Hash received message: h = SHA-256(message)
  4. If h == h': signature valid ✅
```

**Provides:**
- **Authentication**: Proves sender identity
- **Integrity**: Detects tampering
- **Non-repudiation**: Signer cannot deny signing

### HMAC (Hash-based Message Authentication Code)

**Purpose**: Verify both authenticity and integrity using a shared key

**Process:**
```
HMAC = Hash(outer_key + Hash(inner_key + message))
```

**Example (Python):**
```python
import hmac
import hashlib

key = b"shared_secret_key"
message = b"important data"

# Create HMAC
h = hmac.new(key, message, hashlib.sha256)
authentication_tag = h.digest()

# Verify (receiver also computes and compares)
h2 = hmac.new(key, message, hashlib.sha256)
if h.digest() == h2.digest():
    print("Message authentic ✅")
```

### End-to-End Encryption (E2EE)

**Used in:** WhatsApp, Signal, Facebook Messenger

**Signal Protocol Components:**
1. **X3DH Key Agreement**: Initial key exchange
2. **Double Ratchet Algorithm**: Evolving keys per message
3. **Forward Secrecy**: Old messages safe even if current key compromised
4. **Post-Compromise Security**: Future messages safe after key recovery

**Security Properties:**
- ✅ Forward secrecy
- ✅ Post-compromise security
- ✅ Authentication
- ✅ Integrity verification
- ✅ Protection against MITM attacks
- ✅ Deniability (no cryptographic proof)

### Merkle Trees in Blockchain

**Purpose**: Efficiently verify blockchain data integrity

**Structure:**
```
                    Root Hash
                    (H1234)
                    /      \
                  /          \
              H12              H34
              / \              / \
             /   \            /   \
           H1    H2        H3    H4
           |     |         |     |
          TX1   TX2       TX3   TX4
```

**Properties:**
- Any change to transaction changes all hashes above it
- O(log n) verification time
- Complete transaction history in single hash
- Used in Bitcoin and all blockchain systems

---

## Security Threats & Prevention

### Man-in-the-Middle (MITM) Attack

**Attack Vector:**
```
Alice -----> Attacker -----> Server
     (thinks talking to Server)
```

Attacker intercepts and can:
- Eavesdrop on communications
- Modify messages
- Impersonate either party
- Perform session hijacking

**Prevention:**
- ✅ **HTTPS/TLS**: Encrypt all data
- ✅ **Certificate Pinning**: Verify server certificate
- ✅ **MFA**: Multi-factor authentication
- ✅ **DNSSEC**: Secure DNS queries
- ✅ **VPN**: Encrypt entire connection
- ✅ **Public Key Verification**: Verify public keys out-of-band

### DNS Spoofing

**Attack:** Attacker returns false IP address for domain

**Prevention:**
- DNSSEC: Digitally sign DNS records
- DNS over HTTPS (DoH): Encrypt DNS queries
- DNS over TLS: Port 853 with TLS

### Collision Attacks

**Attack:** Find two inputs that hash to same value

**Prevention:**
- Use SHA-256 or SHA-3 (collision-resistant)
- Avoid MD5, SHA-1 for security purposes

### Password Cracking

**Threats:**
- Brute force: Try all combinations
- Dictionary: Common passwords
- Rainbow tables: Precomputed hashes

**Prevention:**
- Use salted hashes: `hash(password + salt)`
- Use key derivation functions: Argon2, PBKDF2
- Enforce strong password policies
- Rate limiting on login attempts

### Nonce Reuse in Encryption

**Risk:** Reusing IV/nonce with same key breaks security

**Prevention:**
- Generate unique nonce for each encryption
- Use larger nonce sizes (XChaCha20: 192 bits)
- Never reuse IV in CBC mode
- Use GCM or AEAD modes properly

---

## Implementation Best Practices

### Cryptographic Key Management

**DO:**
- ✅ Use secure random number generator for keys
- ✅ Store keys in Hardware Security Module (HSM)
- ✅ Rotate keys periodically
- ✅ Use different keys for different purposes
- ✅ Derive keys from passwords using KDF

**DON'T:**
- ❌ Hardcode keys in source code
- ❌ Use weak key derivation (simple hashing)
- ❌ Reuse keys across different applications
- ❌ Store keys in plaintext
- ❌ Use default/weak passwords

**Example (Python - Secure Key Derivation):**
```python
from argon2 import PasswordHasher
import secrets
import os

# Key derivation from password
ph = PasswordHasher()
password = "strong_password_123"
hashed = ph.hash(password)

# Verify
ph.verify(hashed, password)

# Cryptographic key generation
key = secrets.token_bytes(32)  # 256-bit key
```

### Cipher Selection

**For Bulk Data Encryption:**
- Use AES-256-GCM (with hardware support)
- Or ChaCha20-Poly1305 (without hardware support)

**For Key Exchange:**
- Use ECDH with P-256 or Curve25519
- Avoid RSA (slower), use only if required

**For Digital Signatures:**
- Use Ed25519 (elliptic curve, fast)
- Or RSA-4096 (established, slower)

**For Hashing:**
- Use SHA-256 for general purposes
- Use SHA-512 for maximum security
- Never use MD5 or SHA-1

### Protocol Implementation

**TLS/SSL Configuration:**
```nginx
# HTTPS server configuration
ssl_protocols TLSv1.2 TLSv1.3;
ssl_ciphers HIGH:!aNULL:!MD5;
ssl_prefer_server_ciphers on;
ssl_session_timeout 1d;
ssl_session_cache shared:SSL:50m;
ssl_session_tickets off;
```

**SSH Configuration:**
```bash
# ~/.ssh/config
Host *
    HostKeyAlgorithms ssh-ed25519,rsa-sha2-512
    PubkeyAcceptedAlgorithms ssh-ed25519,rsa-sha2-512
    Ciphers chacha20-poly1305@openssh.com,aes-256-gcm@openssh.com
    KexAlgorithms curve25519-sha256,curve25519-sha256@libssh.org
```

### Security Checklist

- [ ] Use TLS 1.2 or higher
- [ ] Use 256-bit symmetric keys
- [ ] Use 2048+ bit RSA or 256-bit ECC
- [ ] Hash passwords with salts and KDF
- [ ] Never hardcode secrets
- [ ] Use HMAC for message authentication
- [ ] Implement key rotation
- [ ] Log security events
- [ ] Keep libraries updated
- [ ] Conduct security audits

---

## Conclusion

Modern cryptography and secure network communication rely on:

1. **Hash Functions** (SHA-256): Data integrity, password storage
2. **Symmetric Encryption** (AES-256-GCM): Fast bulk encryption
3. **Asymmetric Encryption** (ECDH, RSA): Secure key exchange, signatures
4. **Network Protocols** (TLS, SSH): Practical secure communication
5. **Authentication & Integrity** (HMAC, Digital Signatures): Trust establishment

Understanding these fundamentals is essential for building secure systems, implementing strong authentication, protecting data in transit and at rest, and defending against modern cyber threats.

**Remember:** Security is a process, not a product. Stay updated with best practices, keep dependencies current, and regularly audit your systems.

---

## Resources

- [NIST Cryptographic Standards](https://nvlpubs.nist.gov/)
- [OWASP Cryptographic Storage Cheat Sheet](https://cheatsheetseries.owasp.org/)
- [RFC 8446 - TLS 1.3](https://tools.ietf.org/html/rfc8446)
- [Signal Protocol Documentation](https://signal.org/docs/)
- [Practical Cryptography for Developers](https://cryptobook.nakov.com/)

---

## Author Notes

This repository is maintained as an educational resource. Always consult official security guidelines and conduct threat modeling for your specific use cases before deploying cryptographic systems in production.

**Last Updated:** December 2025