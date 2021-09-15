<p align="center">
    <img src="assets/keystone.png" alt="logo" height="200" />
    <h1 align="center">Project Amnesia</h1>
</p>
<p align="center">A Multi-threaded key-value pair store using thread safe locking mechanism allowing concurrent reads.</p>

## Curious to Try it out??
Check out NodeJS client [here](https://github.com/NikhilCodes/amnesia-js).
Shell Client & Web Client Coming soon

## Documentation
### 1. Simple Key Value pair
```bash
IN:  SET key42 AS something
OUT: OK

IN:  GET key42
OUT: something

IN:  SET key42 AS "multi word"
OUT: OK

IN:  GET key42
OUT: multi word
```

### 2. Simple Key Value pair But with Expiry
```bash
IN:  SET key42 AS something WHERE TTL=30s
OUT: OK

IN:  GET key42
OUT: something

# After 30 seconds

IN:  GET key42
OUT: <NIL>
```

### 3. ...With Limited Number of Read Op
```bash
IN:  SET key42 AS something WHERE NFETCH=2
OUT: OK

IN:  GET key42  # First Read Op
OUT: something

IN:  GET key42  # Second Read Op Now Deleting
OUT: something

IN:  GET key42
OUT: <NIL>
```

## TODO
 - [x] Implement Thread Safe CHM
 - [x] NodeJs Client
 - [ ] Shell-based client 
 - [ ] Web Playground 
 - [ ] Multi Tenant

