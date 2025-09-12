import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

export async function login(username: string, password: string) {
  try {
    const response = await axios.post(`${API_URL}/auth/login`, {
      username: username,
      password: password,
    });
    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function getAlerts(lat: number, lng: number, radius: number, token: string) {
  try {
    const response = await axios.get(`${API_URL}/alerts`, {
      params: { lat, lng, radius },
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.data.alerts;
  } catch (error) {
    console.error("Failed to fetch alerts:", error);
    throw error;
  }
}