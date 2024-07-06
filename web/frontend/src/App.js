import React, { useState } from 'react';
import './App.css';

function App() {
  const [organisationName, setOrganisationName] = useState('');
  const [results, setResults] = useState([]);

  const handleSearch = async () => {
    try {
      const response = await fetch(`http://localhost:8080/organisation/${organisationName}`);
      const data = await response.json();

      if (Array.isArray(data)) {
        setResults(data);
      } else {
        setResults([]);
      }
    } catch (error) {
      console.error('Error fetching data:', error);
      setResults([]);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Does It Sponsor?</h1>
        <div className="search-bar">
          <label htmlFor="organisation-name">Organisation Name</label>
          <input
            id="organisation-name"
            type="text"
            value={organisationName}
            onChange={(e) => setOrganisationName(e.target.value)}
          />
          <button onClick={handleSearch}>Search</button>
        </div>
        <div className="results-box">
          {results.length > 0 ? (
            <table>
              <thead>
                <tr>
                  <th>Organisation Name</th>
                  <th>City</th>
                  <th>County</th>
                  <th>Type</th>
                  <th>Route</th>
                </tr>
              </thead>
              <tbody>
                {results.map((org, index) => (
                  <tr key={index}>
                    <td>{org.organisation_name}</td>
                    <td>{org.city}</td>
                    <td>{org.county}</td>
                    <td>{org.type}</td>
                    <td>{org.route}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          ) : (
            <p>No results found</p>
          )}
        </div>
      </header>
    </div>
  );
}
export default App;
