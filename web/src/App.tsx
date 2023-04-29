import './App.css'

type Item = {
  name: string;
  data: string;
};

type Repository = {
  repository: Item[];
};

async function getRepository() {
  try {
      const response = await fetch(
          'http://localhost:8080/feed',
          {   
              method: 'GET',
              headers: {
                  Accept : 'application/json'
              }
          },
      );

      if (!response.ok) {
          throw new Error(`error, status: ${response.status}`);
      }

      const result = (await response.json());
      console.log(JSON.stringify(result, null, 4));
  } catch (error) {
      if (error instanceof Error) {
          console.log('error message: ', error.message);
      } else {
          console.log('unexpected error: ', error);
      }
  }
}

function App() {
  return (
    <>
      <div className="card">
        <button onClick={ async () => { await getRepository(); } }>
          test get 
        </button>
      </div>
    </>
  )
}

export default App
