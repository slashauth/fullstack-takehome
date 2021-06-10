# Debrief Stonks
This package is built of two separate components -- backend and frontend. Each package is described below

## Backend
This server is a very simple scaffolding of a GQL server that you can use.

We have stubbed out some methods and most GQL resolvers will panic so please implement them!

### Docker Compose
There is a docker compose file within the project and will create a composed stck consisting of this API as well as a Redis server. You can run it by executing

`docker-compose up --build`

NOTE: Without the `--build` the service will start but will not rebuild so no changes will show up.

The server will then be available on `http://localhost:8080`. You can explore the existing schema by navigating to `http://localhost:8080/graphiql`.

### Storage and Persistence
While Redis is not required, it's quite easy to setup and use so we included it in this docker-compose file. Feel free to rip it out and use something else like postgres, mongo, or really anything you prefer!

Keep in mind that if you choose to use Redis, you will need to run `redis-server` on your machine if you plan to run the API locally. It looks for a server running on port 6379 by default.

### Generating the schema
When you make updates to the schema file `schema.graphqls` you can regenerate the code by running `make schema` within this directory.

## Frontend
The frontend is generated via create-react-app and uses apollo to generate gql stubs. `codegen.yml` has this configuration.

### Generating schema
If you change the schema on the backend or the frontend and need to regenerate, you can run `yarn run generate` from within the frontend directory. Note that your backend must be running on port 8080 for this to work.

### Running the app
Simply run `yarn start` and the app will start. It communicates with a backend on `localhost:8080` so ensure the backend is up and running. The app will automatically reload with any changes.

### Adding new packages
Use `yarn add` to add new packages.