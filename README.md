# smart-game
Smart game solver<br>
Для первой задачи на Golang решил оцифровать такую настольную головоломку:
![Smart game example](https://cdn.mosigra.ru/mosigra.product.main/540/771/DSC_2908_800x500.jpg)
<br>
Получилось как-то так:
![Screenshot](https://github.com/demisang/smart-game/raw/master/screenshot.png)
<br>
Запускаем, открываем http://localhost:8770/ и можно мышкой выбирать какую фигуру двигать/вращать, можно на клаве.<br>
Но ещё не рализовано удаление фигуры, размещение фигуры и самое важное - перебрать все возможные комбинации
из свободных фигур, чтобы доска полностью заполнилась.