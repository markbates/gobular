require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.js");

$(() => {
  $("#checker-Expression").on("change keyup paste", (e) => {
    $("#checker-form").submit();
  });
  $("#checker-TestString").on("change keyup paste", (e) => {
    $("#checker-form").submit();
  });

  $("#checker-form").on("submit", (e) => {
    e.preventDefault();
    let f = $(e.target);
    let sd = f.serialize();
    history.pushState({}, "", `/check?${sd}`);
    $.post(f.attr("action"), sd, (data) => {
      $("#results").replaceWith(data);
    });
  });
});
