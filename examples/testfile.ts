export const locationWorkerSettings = {
  locationWorkerlogs: false,
  debugLocationWorker: false,
  debugPermissions: false,
  debugInitialization: false,
};

export const databaseSettings = {
  environment: "development",
  database: {
    name: "databaseName",
    location: "default",

    //  \/  \/  \/  \/  \/  \/

    resetDatabase: false,
    resetAsyncStorage: false,

    //  /\  /\  /\  /\  /\  /\

    fakeDataCountWhenEmptyTodos: 0,
    fakeDataCountWhenEmptyTags: 0,
  },
  logging: {
    databaseProviderLogs: false,
    todoCreationLogs: false,
    todoGetLogs: false,
    todoUpdateLogs: false,
    todoDeleteLogs: false,
    todoNotificationLogs: false,
    tagCreationLogs: false,
    tagGetLogs: false,
    tagUpdateLogs: false,
    tagDeleteLogs: false,
    todoTagCreationLogs: false,
    todoTagGetLogs: false,
    todoTagRemovalLogs: false,
    statsLogs: false,
    syncLogs: false,
  },
};
