import { useRef, useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'

function App() {
  // Ref for query
  const queryRef = useRef(null);
  const [data, setData] = useState({});
  const [response, setResponse] = useState(null);
  const [glowPrimary, setGlowPrimary] = useState(false);
  const [glowSecondary, setGlowSecondary] = useState([false, false, false, false]);

  const blipPrimary = () => {
    setGlowPrimary(true);
    setTimeout(() => setGlowPrimary(false), 1000);
  }

  const blipSecondary = (index) => {
    console.log(index)
    const newGlowSecondary = [...glowSecondary];
    newGlowSecondary[index] = true;
    setGlowSecondary([...newGlowSecondary]);
    newGlowSecondary[index] = false;
    setTimeout(() => setGlowSecondary([...newGlowSecondary]), 1000);
  }

  return (
    <div className="App">
      <div className={'left-half'}>
        <input
          className={'query'}
          ref={queryRef}
          onKeyPress={(e) => {
            if (e.key === 'Enter') {
              fetch(`http://localhost:4422/cmd`, {
                method: 'POST',
                body: JSON.stringify({ query: queryRef.current.value }),
                headers: {
                  'Content-Type': 'application/json',
                }
              })
                .then(res => res.json())
                .then((data) => {
                  setResponse(JSON.stringify(data, null, 2));
                  if (queryRef.current.value.startsWith('SET ')) {
                    const varName = queryRef.current.value.split(' ')[1].replace(/"/g, '');
                    const value = queryRef.current.value.split(' ')[3].replace(/"/g, '');
                    setData(d => ({ ...d, [varName]: value }))
                    blipPrimary();
                  } else {
                    blipSecondary(Math.floor(Math.random() * 4));
                  }
                });
            }
          }}
        />
        <div style={{ fontSize: 20, fontWeight: 'bold' }}>Insert History</div>
        <div className={'code'}>
          {JSON.stringify(data, null, 2).split(/\n/).map(line => <div key={line}>{line}</div>)}
        </div>
        <br/>
        <div style={{ fontSize: 20, fontWeight: 'bold' }}>Response</div>
        <div className={'code'}>
          {response}
        </div>
      </div>
      <div className={'right-half'}>
        <div style={{ display: 'flex', justifyContent: 'center' }}>
          <div className={'card'} style={{ background: glowPrimary ? 'yellow' : 'white' }}>
            <img src="https://img.icons8.com/dusk/64/null/database.png" alt={'primary-db'}/>
          </div>
        </div>

        <div style={{ display: 'flex', justifyContent: 'center' }}>
          <div className={'card small'} style={{ background: glowSecondary[0] ? 'yellow' : 'white' }}>
            <img src="https://img.icons8.com/dusk/64/null/database.png" alt={'secondary-db'}/>
          </div>
          <div className={'card small'} style={{ background: glowSecondary[1] ? 'yellow' : 'white' }}>
            <img src="https://img.icons8.com/dusk/64/null/database.png" alt={'secondary-db'}/>
          </div>
          <div className={'card small'} style={{ background: glowSecondary[2] ? 'yellow' : 'white' }}>
            <img src="https://img.icons8.com/dusk/64/null/database.png" alt={'secondary-db'}/>
          </div>
          <div className={'card small'} style={{ background: glowSecondary[3] ? 'yellow' : 'white' }}>
            <img src="https://img.icons8.com/dusk/64/null/database.png" alt={'secondary-db'}/>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
