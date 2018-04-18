require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

$(() => {
  $("#expression-Expression").on("change keyup paste", () => {
    $("#expression-form").submit();
  });

  $("#expression-TestString").on("change keyup paste", () => {
    $("#expression-form").submit();
  });

  $("#expression-form").on("submit", (e) => {
    e.preventDefault();
    let f = $(e.target);
    let sd = f.serialize();
    $.post(f.attr("action"), sd, (data) => {
      $("#results").replaceWith(data);
      let href = $("#rerun-url").attr("href");
      history.pushState({}, "", href);
    });
  });
});
