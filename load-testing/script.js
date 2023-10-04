import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
    http.get('https://test.k6.io');
    sleep(1);
}

// k6 run --vus 10 --duration 1m .\load-testing\script.js