import React, { useState } from 'react';
import { createRoot } from 'react-dom/client';
import './style.css';


const App = () => {
  return (
    <>
      <div>
        <h1>Hello World!</h1>
      </div>
    </>
  );
}

createRoot(document.getElementById('app')).render(<App />);
