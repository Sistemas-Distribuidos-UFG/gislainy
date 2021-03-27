
import Event from "events";
import { Socket } from "net"

export default class SocketClient {

    #serverConnection = {}
    #serverListener = new Event();

    constructor({ host, port }) {
        this.port = port;
        this.host = host
    }

    sendMessage (event, message) {
        this.#serverConnection.write(JSON.stringify({ event, message }))
    }

    attachEvents (events) {
        this.#serverConnection.on('data', data => {
            try {
                data
                    .toString()
                    .split('\n')
                    .filter(line => !!line)
                    .map(JSON.parse)
                    .map(({ event, message }) => {
                        console.log({event, message})
                        this.#serverListener.emit(event, message)
                    })
            } catch (error) {
                console.error('invalid!', data.toString(), error)
            }
        });

        this.#serverConnection.on('end', () => {
            console.log('I disconnected!!')
        })
        this.#serverConnection.on('error', (error) => {
            console.error('Crash server', error)
        })

        for (const [key, value] of events) {
            this.#serverListener.on(key, value);
        }
    }

    async #createConnection () {
        const client = new Socket();

        return new Promise(resolve => {
            client.connect(this.port, this.host, () => {
                resolve(client)
            });
        });
    
    }

    async initialize () {
        this.#serverConnection = await this.#createConnection();
        console.log('I connected to the server!!')
    }


}