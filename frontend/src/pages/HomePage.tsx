import { useState, useEffect } from 'react';
import { getAlerts } from '../services/api';

interface Alert {
  id: string;
  title: string;
  description: string;
}

export function HomePage() {
  const [alerts, setAlerts] = useState<Alert[]>([]);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchAlertsForUserLocation() {
      const token = localStorage.getItem('authToken');
      if (!token) {
        setError("Você não está autenticado.");
        setLoading(false);
        return;
      }

     navigator.geolocation.getCurrentPosition(
        async (position) => {
          const { latitude, longitude } = position.coords;
          console.log(`Localização obtida: Lat ${latitude}, Lng ${longitude}`);
          
          try {
            const searchRadius = 5000;
            const data = await getAlerts(latitude, longitude, searchRadius, token);
            setAlerts(data || []);
          } catch (err) {
            setError("Falha ao buscar alertas.");
          } finally {
            setLoading(false);
          }
        },
        (geoError) => {
          console.error("Erro ao obter geolocalização:", geoError);
          setError("Não foi possível obter sua localização. Por favor, habilite a permissão no seu navegador.");
          setLoading(false);
        }
      );
    }

    fetchAlertsForUserLocation();
  }, []);

  return (
    <div>
      <h2>Alertas Próximos</h2>
      {loading && <p>Obtendo sua localização e buscando alertas...</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {!loading && !error && alerts.length === 0 && <p>Nenhum alerta encontrado perto de você.</p>}
      <ul>
        {alerts.map(alert => (
          <li key={alert.id}>
            <strong>{alert.title}</strong>: {alert.description}
          </li>
        ))}
      </ul>
    </div>
  );
}