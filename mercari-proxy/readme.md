# Proxy Service

## API Spec
### GetProvince
#### Request
`GET /api/provinces`
#### Response
```
[
    {
        "id": "1",
        "name": "Jawa Tengah"
    }
]
```

### GetCities
#### Request
`GET /api/cities`
#### Response
```
[
    {
        "id": "1",
        "name": "Semarang",
        "province_name": "Jawa Tengah"
    },
    {
        "id": "2",
        "name": "Surakarta",
        "province_name": "Jawa Tengah"
    }
]
```

### GetDistricts
#### Request
`GET /api/districts`
#### Response
```
[
    {
        "id": "1",
        "name": "Banyumanik",
        "city_name": "Semarang"
    },
    {
        "id": "2",
        "name": "Banjarsari",
        "city_name": "Surakarta"
    }
]
```

### GetSubDistricts
#### Request
`GET /api/subdistricts`
#### Response
```
[
    {
        "id": "1",
        "name": "Gedawang",
        "district_name": "Banyumanik"
    },
    {
        "id": "1",
        "name": "Banyuanyar",
        "district_name": "Banjarsari"
    }
]
```

## Cache
The cache is invalidated automatically after 24 hours after each API is called. The drawback is that if there is any new data changes, it might need some time until the user see the latest data

## Scaling
Since this service is a proxy and will use cache a lot, doing vertical scaling by having a bigger redis capacity or implementing redis cluster for redis horizontal scaling might help.

For the service itself, it can be scaled horizontally by adding the number of instances or vertically by having a higher per-instance capacity