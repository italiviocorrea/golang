window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">
  hostname = window.location.hostname;
  port = window.location.port;
  if (hostname === 'localhost') {
      port = '8080';
  }	
  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: window.location.protocol+"//"+hostname+":"+port+"/api/v1/swagger/swagger3.json",
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
