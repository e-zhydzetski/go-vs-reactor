FROM gradle:6.0.1-jdk8 AS build
WORKDIR /workspace
COPY . .
RUN gradle bootJar --no-daemon

FROM openjdk:8u222-jre-slim
RUN apt-get update && apt-get install -y dos2unix
WORKDIR /app
COPY ./launch-jar.sh ./launch-jar.sh
RUN dos2unix ./launch-jar.sh
EXPOSE 8080
ENV JAVA_OPTS="-XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap"
COPY --from=build workspace/build/libs/*.jar /app/service.jar
ENTRYPOINT ["./launch-jar.sh","service.jar"]