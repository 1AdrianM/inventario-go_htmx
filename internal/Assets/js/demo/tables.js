document.body.addEventListener('htmx:afterSwap', function(event) {
    if (event.target.id === 'dataTable') {
        const modal = new bootstrap.Modal(document.getElementById('user_modal'));
        modal.show()

    }
});
document.body.addEventListener('htmx:afterSwap', function(event) {
    if (event.target.id === 'dataTable') {
        const umodal = new bootstrap.Modal(document.getElementById('Updateuser_modal'));
        umodal.show()

    }
});

document.body.addEventListener('htmx:afterSwap', function(event) {
    if (event.target.id === 'Import_modal') {
        const umodal = new bootstrap.Modal(document.getElementById('Import_modal'));
        umodal.show()

    }
});
