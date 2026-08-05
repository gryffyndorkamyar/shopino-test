# Shopino Svelte Frontend

صفحه اصلی و صفحات مرتبط شاپینو با **SvelteKit**، متصل به Gateway روی پورت `8088`.

## صفحات

| مسیر | API |
|------|-----|
| `/` | `GET /api/products/` |
| `/products` | `GET /api/products/` |
| `/products/[id]` | `GET /api/products/:id` |
| `/stores` | `GET /api/stores/` |
| `/login` `/register` | Django auth |

## اجرا

```powershell
docker compose up -d
cd frontend\svelte
npm install
npm run dev
```

باز کن: http://127.0.0.1:5173

## حساب دمو

- username: `demo`
- password: `demo12345`
