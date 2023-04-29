import fetch from 'node-fetch';
import { Repository } from '../types/repository';

export async function getRepository() {
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

      const result = (await response.json()) as Repository;
      console.log(JSON.stringify(result, null, 4));
  } catch (error) {
      if (error instanceof Error) {
          console.log('error message: ', error.message);
      } else {
          console.log('unexpected error: ', error);
      }
  }
}

export function print() {
    console.log('hellope')
}
