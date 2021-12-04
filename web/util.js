const CooterParser = (response) => {
    const parsed = response['passed'].map(rpm => JSON.parse(rpm));
    const errored = response['error'].map(rpm => JSON.parse(rpm));
    return [parsed, errored];
};