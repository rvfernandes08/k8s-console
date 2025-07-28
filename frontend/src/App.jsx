import { useEffect, useState } from 'react';

function App() {
  const [namespaces, setNamespaces] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchNamespaces = async () => {
    setLoading(true);
    try {
      const res = await fetch('/api/namespaces');
      const data = await res.json();
      setNamespaces(data.namespaces || []);
    } catch (err) {
      console.error('Erro ao buscar namespaces:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchNamespaces();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100 p-4">
      <h1 className="text-2xl font-bold mb-4">Painel Kubernetes</h1>
      <button
        className="mb-4 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
        onClick={fetchNamespaces}
      >
        Atualizar Namespaces
      </button>
      {loading ? (
        <p>Carregando...</p>
      ) : (
        <ul className="list-disc pl-6">
          {namespaces.map(ns => (
            <li key={ns}>{ns}</li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default App;

