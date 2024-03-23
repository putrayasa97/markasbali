# Step untuk menjalankan project ini
- Setup koneksi database pada file .env
- --DB_HOST = 127.0.0.1
- --DB_USER = root
- --DB_PASS = root
- --DB_PORT = 8889
- --DB_NAME = perpus // untuk database aplikasi
- --DB_NAME_TESTING = perpus_testing // untuk database testing
- Membuat database default untuk aplikasi dan testing
# Step untuk menjalankan testing
- > go test -v -count=1 ./...

