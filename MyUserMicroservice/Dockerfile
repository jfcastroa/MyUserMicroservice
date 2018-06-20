FROM microsoft/aspnetcore:2.0 AS base
WORKDIR /app
EXPOSE 80

FROM microsoft/aspnetcore-build:2.0 AS build
WORKDIR /src
COPY MyUserMicroservice/MyUserMicroservice.csproj MyUserMicroservice/
RUN dotnet restore MyUserMicroservice/MyUserMicroservice.csproj
COPY . .
WORKDIR /src/MyUserMicroservice
RUN dotnet build MyUserMicroservice.csproj -c Release -o /app

FROM build AS publish
RUN dotnet publish MyUserMicroservice.csproj -c Release -o /app

FROM base AS final
WORKDIR /app
COPY --from=publish /app .
ENTRYPOINT ["dotnet", "MyUserMicroservice.dll"]
