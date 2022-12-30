export default new (class API {
  public Get = async (url: string): Promise<any> => {
    // `${window.location.protocol}//${window.location.host}${url}`
    return await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res) => {
      return res.json();
    });
  };
})();
