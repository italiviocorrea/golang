window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">
  hostname = window.location.hostname;
  if (hostname === 'localhost') {
      hostname += ':8080';
  }	
  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: window.location.protocol+"//"+hostname+"/api/v1/swagger/swagger3.json",
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
