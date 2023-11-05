import http from "k6/http";
import { check, sleep } from "k6";

const BASE_URL = __ENV.FRONTEND_URL || 'http://localhost:3333';

export const options = {
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(99)<1000'],
  },
  scenarios: {
    load: {
      executor: 'constant-arrival-rate',
      rate:  10,
      preAllocatedVUs: 5,
      maxVUs: 20,
      exec: 'requestPizza',
      startTime: '0s',
      duration: '30s',
    }
  }
};

export function setup() {
  let res = http.get(BASE_URL)
  if (res.status !== 200) {
    throw new Error(`Got unexpected status code ${res.status} when trying to setup. Exiting.`)
  }
}

export function requestPizza () {
  let restrictions = {
    maxCaloriesPerSlice: 500,
    mustBeVegetarian: false,
    maxNumberOfToppings: 6,
    minNumberOfToppings: 2
  }
  let res = http.post(`${BASE_URL}/api/pizza`, JSON.stringify(restrictions), {
    headers: {
      'Content-Type': 'application/json',
      'X-User-ID': 23423,
    },
  });
}
