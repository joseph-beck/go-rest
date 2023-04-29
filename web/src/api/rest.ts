import fetch from 'node-fetch';
import { Repository } from '../types/repository';
import { error } from 'console';

export async function get() {
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
        return JSON.stringify(result, null, 4);
    } catch (error) {
        if (error instanceof Error) {
            console.log('error message: ', error.message);
        } else {
            console.log('unexpected error: ', error);
        }
    }
}

export async function post(name: string, data: string) {
    try {
        const response = await fetch(
            'http://localhost:8080/feed',
            {
                method: 'POST',
                headers: {
                    'content-type' : 'application/json' 
                },
                body : JSON.stringify({
                    name : name,
                    data: data
                }),
            }
        );

        if (!response.ok) {
            throw new Error(`error, status: ${response.status}`);
        }
    } catch (error) {
        if (error instanceof Error) {
            console.log('error message: ', error.message);
        } else {
            console.log('unexpected error: ', error);
        }
    }
}