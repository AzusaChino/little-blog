interface Config {
  url?: string
}

const prod: Config = {
  url: 'http://119.45.30.109:8080'
}

const dev: Config = {
  url: 'http://127.0.0.1:8080'
}

export const config = process.env.NODE_ENV === 'development' ? dev : prod;
