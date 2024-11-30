import React, { useState } from 'react';
import './App.css';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e, endpoint) => {
    e.preventDefault();
    setError('');

    const data = { username, password };

    try {
      const response = await fetch(`http://localhost:8080/${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        alert('Успешно!');
      } else {
        setError('Ошибка при отправке данных');
      }
    } catch (err) {
      setError('Ошибка при соединении с сервером');
    }
  };

  return (
    <div className="container">
      <h1 className="header">Регистрация</h1>
      <div className="form-container">
        <form onSubmit={(e) => handleSubmit(e, 'signup')}>
          <div className="input-group">
            <label>Username</label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="input-group">
            <label>Password</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="button">Авторизация</button>
        </form>
      </div>
      {error && <p className="error">{error}</p>}
    </div>
  );
}

export default App;
