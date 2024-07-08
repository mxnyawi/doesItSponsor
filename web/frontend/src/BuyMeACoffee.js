import React from 'react';

function Coffee() {
  return (
    <a
      className="buyButton"
      target="_blank"
      rel="noopener noreferrer"
      href="https://www.buymeacoffee.com/mxnyawi"
    >
      <img
        className="coffeeImage"
        src="https://cdn.buymeacoffee.com/buttons/bmc-new-btn-logo.svg"
        alt="Buy me a coffee"
      />
      <span className="coffeeButtonText">Buy me a coffee :)</span>
    </a>
  );
}

export default Coffee;
