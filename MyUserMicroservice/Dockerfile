FROM microsoft/dotnet:2.1-aspnetcore-runtime AS base
WORKDIR /app
EXPOSE 80

FROM microsoft/dotnet:2.1-sdk AS build
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
