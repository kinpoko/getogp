# getogp

API for getting ogp written in go with Vercel

## Usage

```bash
curl https://getogp.vercel.app/api?url=https://github.com/kinpoko/getogp | jq .
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   387  100   387    0     0    778      0 --:--:-- --:--:-- --:--:--   778
{
  "title": "GitHub - kinpoko/getogp: API for getting ogp written in go with Vercel",
  "url": "https://github.com/kinpoko/getogp",
  "description": "API for getting ogp written in go with Vercel. Contribute to kinpoko/getogp development by creating an account on GitHub.",
  "image": "https://opengraph.githubassets.com/a9d3d700b2e73e85a0dee8fad2d1963dfff28d2fb8190d8bae0fa9e77e1f75e7/kinpoko/getogp"
}
```
