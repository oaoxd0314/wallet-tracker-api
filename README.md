# Wallet Tracker Api

>使用 GO + GIN 打造的 API Server

## usage

下載專案

```
git clone https://github.com/oaoxd0314/wallet-tracker-api.git
```

在跟目錄設置 `.env` 檔案 並寫入以下變數

```
PORT=8080(你想運行的 port)
API_KEY=你的 etherscan api token
```

之後運行 `go run .` 啟動專案

## use Heroku

[下載 Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)

登入 heroku

```
hreoku login
```

推送上 heroku
```
git push heroku
```

開啟 heroku App
```
heroku open
```
