import React from 'react';
import './App.css';
// import logo from './logo.png';
// import TugasIntroReact from './Tugas-Intro-ReactJS/TugasIntroReact';
// import TugasHooks from './Tugas-Hooks/TugasHooks';
// import TugasCRUDHooks from './TugasCRUDHooks/TugasCRUDHooks';

import TugasAxios from './Tugas-Axios/TugasAxios';

function App() {
    return (
        <div className="App">
            {/* <TugasHooks />
      <div className="container">
        <img src={logo} alt="Sanbercode Logo" className="logo" />
        <h1>THINGS TO DO</h1>
        <p>During bootcamp in sanbercode</p>
        <TugasIntroReact />
      </div> */}
            {/*<TugasCRUDHooks />*/}
            <TugasAxios />
        </div>
    );
}

export default App;
