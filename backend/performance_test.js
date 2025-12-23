import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  vus: 10, 
  duration: '10s', 
};

export default function () {
  const res = http.get('http://localhost:3000/');
  
  check(res, {
    'durum kodu 200 mü?': (r) => r.status === 200,
    'cevap süresi 200ms altında mı?': (r) => r.timings.duration < 200,
  });

  sleep(1);
}