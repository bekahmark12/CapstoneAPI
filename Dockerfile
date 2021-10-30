FROM public.ecr.aws/lambda/go:1.x
COPY [., /src/] ${LAMBDA_TASK_ROOT}
# Copy function code
# COPY connectapi/userservice/handlers/middleware.go ${LAMBDA_TASK_ROOT}
# COPY connectapi/userdb/config/init.sh ./
# COPY connectapi/userservice/auth/jwt.go ./
# COPY connectapi/userservice/data ./
WORKDIR "/src"
COPY . .
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "app.lambda_handler" ]


ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV APP_DB_USER=admin
ENV APP_DB_PASS=password
ENV USER_DB_NAME=users
ENV APP_DB_PORT=5432

# API_SERVICE_PORT=8081

# MONGO_PORT=27017-27019

# SECRET_KEY=capstonesecretkey123
