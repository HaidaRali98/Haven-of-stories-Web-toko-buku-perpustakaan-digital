import http.server
import socketserver
import webbrowser

PORT = 8000
Handler = http.server.SimpleHTTPRequestHandler

print(f"🚀 Server berjalan di http://localhost:{PORT}")
print("Membuka browser otomatis...")

webbrowser.open(f'http://localhost:{PORT}')

with socketserver.TCPServer(("", PORT), Handler) as httpd:
    print("Tekan CTRL+C untuk berhenti.")
    httpd.serve_forever()