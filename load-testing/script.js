import http from 'k6/http';
import { check, group } from 'k6';

const API_ENDPOINT = 'http://13.127.57.44:8080/user/login'; // Replace with your API endpoint
const USERS_DATA_FILE = 'users_data.json'; // Replace with the path to your JSON file

// Load user data from a JSON file
const usersData = JSON.parse(open(USERS_DATA_FILE));

export let options = {
    stages: [
        { duration: '5s', target: 1000 }, // Ramp-up to 1,000 users in 5 seconds
        { duration: '10s', target: 1000 }, // Stay at 1,000 users for 1 minute
    ],
};

export default function () {
    // Randomly select a user from the loaded data
    let randomUser = usersData[Math.floor(Math.random() * usersData.length)];
    let payload = JSON.stringify({
        user_name: randomUser.user_name,
        password: randomUser.password,
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    group('Login API', () => {
        let response = http.post(API_ENDPOINT, payload, params);

        // Check if the response status code is 200
        check(response, {
            'is login successful': (res) => res.status === 200,
        });
    });
}
