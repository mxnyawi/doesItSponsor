import React, { useState, useEffect } from 'react';
import './App.css';
import BuyMeACoffee from './BuyMeACoffee';

function App() {
  const [organisationName, setOrganisationName] = useState('');
  const [results, setResults] = useState([]);
  const [showResults, setShowResults] = useState(false);

  const handleSearch = async () => {
    try {
      const response = await fetch(`http://localhost:8080/organisation/${organisationName}`);
      const data = await response.json();

      if (Array.isArray(data)) {
        setResults(data);
        setShowResults(true);
      } else {
        setResults([]);
        setShowResults(false);
      }
    } catch (error) {
      console.error('Error fetching data:', error);
      setResults([]);
      setShowResults(false);
    }
  };

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      event.preventDefault(); // Prevent default behavior of form submission
      handleSearch();
    }
  };

  useEffect(() => {
    setShowResults(results.length > 0);
  }, [results]);

  return (
    <div className="App">
      <header className="App-header">
        <h1 className="App-title">Does It Sponsor?</h1>
        <p className="App-description">Find out if an organisation sponsors various Visa types.</p>
        <form className="search-bar" onSubmit={(e) => e.preventDefault()}>
          {/* Using form onSubmit to handle Enter key */}
          <input
            id="organisation-name"
            type="text"
            value={organisationName}
            placeholder="Enter organisation name"
            onChange={(e) => setOrganisationName(e.target.value)}
            onKeyDown={handleKeyDown}
          />
          <button type="button" onClick={handleSearch}>Search</button>
        </form>
        {showResults && (
          <div className="results-box">
            {results.length > 0 ? (
              <table className="results-table">
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
        )}

        <BuyMeACoffee />
      </header>
    </div>
  );
}

export default App;
