CREATE TABLE `task` (
  `task_id` int(10) PRIMARY KEY,
  `task` varchar(250) NOT NULL,
  `status` varchar(30) NOT NULL
);

INSERT INTO `task` VALUES
  (1, 'Read an article on React.js', 'Done'),
  (2, 'Organize a meeting', 'Pending');

ALTER TABLE `task`
  MODIFY `task_id` int(10) AUTO_INCREMENT, AUTO_INCREMENT=3;
