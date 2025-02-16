$(document).ready(function() {
    // 채팅방 생성 버튼 클릭 이벤트
    $('.create-room-btn').on('click', function() {
        $.ajax({
            url: '/rooms',
            method: 'POST',
            dataType: 'json',
            success: function(data) {
                if (data.id) {
                    window.location.href = '/chat/' + data.id;
                }
            },
            error: function(xhr, status, error) {
                console.error('채팅방 생성 실패:', error);
                alert('채팅방 생성에 실패했습니다.');
            }
        });
    });

    function loadRoomList() {
        $.ajax({
            url: '/rooms',
            method: 'GET',
            dataType: 'json',
            success: function(data) {
                data.forEach(function(room) {
                    $('.room-list').append(`
                        <div class="room-card" data-room-id="${room.id}">
                            ${room.id}
                        </div>
                    `);
                });
            },
            error: function(xhr, status, error) {
                console.error('채팅방 목록 로드 실패:', error);
                $('.room-list').html('<p>채팅방 목록을 불러오는데 실패했습니다.</p>');
            }
        });
    }

    loadRoomList();
});