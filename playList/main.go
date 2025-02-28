/*
    party play list
    1. attendance
        i. like songs
    2. get next song according to attendance and their liked songs

    clarification:
        1. one attendance -> many liked songs
        2. multiple attendances
        3. select next play song based on various ways
        4. don't need to handle race condition


设计一个 Party Playlist 的功能，开 Party 的时候有 attendees 和他们喜欢的歌曲。要根据 attendees 和他们喜欢的歌曲来决定下一首要播放的歌。

给一个 API，getLikedSongs(String user) 获取某个 attendees 的所有 liked song。

要求实现三个函数：

addAttendee(String user)
removeAttendee(String user)
getNextSong()
问了 getNextSong 怎么被触发，tie 怎么处理。
设计的时候用一个 hashmap 来统计 song 和喜欢这个 song 的人数的关系。一个 set 存现有用户。这样的话加用户和 remove 用户是 O(1)，getNextSong 是 O(n)。

Follow-up 说如果 song 特别多，优化 getNextSong()。用堆，但是用堆有一个比较复杂的点。Java 里的 PriorityQueue 不太好删除，得遍历，所以要用 lazy update。调用 getNextSong 的时候去 check，这样只需要检查堆顶的元素是不是过期了，摒弃下来的时间是 O(1)。
*/

type song struct {
    id int64
    playTime time.Time
    popularity int
}

type songManager struct {
    songSet map[int64]song
    songMinHeap []songHeap
}

func (s *songManger) addAttendee(userID string) {
    allSongs := getLikedSongs(userID)
}










import (
    "time"
)

type MusicPlayer interface {
    Play(Song)
}

type Song interface {
    GetFilePath() string
    GetPreviousPlayedTime() time.Time
    SetPreviousPlayedTime(time.Time)
}

type Attendance interface {
    GetLikedSongs() []Song
}

type PartyGroup interface {
    GetAttendances() []Attendance
    AddAttendances(Attendance)
    RemoveAttendance(Attendance)
}

type MusicManager interface {
    GetNextSong() Song
    AddSongs([]Song)
}

// Song implementaion
type song struct {
    filePath string
    previousPlayedtime time.Time
}

func (s *song) GetFilePath() string {

}

// PartyGroup
type partyGroup struct {
    attendences []Attendance
    popularityMusicManager MusicManager
}

func (p *partyGroup) GetAttendances() []Attendance {
    return p.attendances
}

func (p *partyGroup) AddAttendance(newAttendance Attandance)  {
    popularityMusicManager.AddSongs(newAttendance.GetLikedSongs())
    p.attendences = append(p.attendences, newAttendance)
}

// Song Min Heap

// MusicManager implementation
type popularityMusicManager struct {
    playedSongs minHeap
    notPlayedSongs []Song
}

/*
    c a b
    ^

    

    1. all music is played once
        i. like count-- when user left
        ii. remove the music if like count == 0
    time complexity: O(1)
    2. not all music is played once (sort all not played music by popularity again)
        i. like count-- when user left
        ii. remove the music if like count == 0
    time complexity: O(mlongm) where m is the not-played song


    choose next song:
        1. previouse play time
        2. popularity
    

*/
func (p *popularityMusicManager) GetNextSong() Song {
    var nextSong Song
    if len(p.notPlayedSongs) != 0 {
        nextSong = p.notPlayedSongs[0]
        p.notPlayedSongs = p.notPlayedSongs[1:]
    } else {
        nextSong = heap.Pop(p.plyaedSongs).(Song)
    }
    nextSong.SetPreviousPlayedTime(time.Now())
    heap.Push(p.playedSongs, nextSong)
    return nextSong
}

func (p *popularityMusicManager) AddSong(newSongs []Song) {
    for _, newSong := range newSongs {
        // TODO: if newSong is in non play songs, add popularity, and sort it again
        // if newSong is in played songs, add popularity
    }
}

func NewPopularityMusicManager(partyGroup PartyGroup) MusicManager {
    popularityCount := make(map[string]int)
    for _, attendance := range partyGroup.GetAttendances() {
        for _, likedSong := range 
    }
}

func main() {
    fmt.Printf("Hello LeetCoder")
}



/*

	what's the most challenging part of amazon, how can i prepare

	what's special about amazon

	
*/