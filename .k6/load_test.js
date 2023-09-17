import http from 'k6/http'
import { check, sleep } from 'k6'

export const options = {
    vus: 400,
    duration: '10s'
}

const params = {
    headers: {
        'Content-Type': 'application/json'
    }
}

export default function() {
    const payload = JSON.stringify({
        take: Math.floor(Math.random() * 3)
    })

    const res = http.post(`${__ENV.URL}`, payload, params)
    check(res, {
        'is status 200': (r) => r.status === 200,
        'valid body': (r) => {
            const obj = JSON.parse(r.body)
            return Object.hasOwn(obj, 'robotTake')
                && Object.hasOwn(obj, 'robotResult')
                && Object.hasOwn(obj, 'userTake')
                && Object.hasOwn(obj, 'userResult')
        },
        'latency': (r) => r.timings.duration <= 60
    })
}