transitionMatrix = eye(14) * 0;
idx = -1;
for i = 1:4
  for j = 1:4
    idx++;
    if i == 1 && j == 1
      continue
    elseif i == 4 && j == 4
      continue
    end
    if j > 1 && idx != 1
      transitionMatrix(idx,idx-1) = 0.25;
    end
    if i > 1 && idx != 4
      transitionMatrix(idx,idx-4) = 0.25;
    end
    if j < 4 && idx != 14
      transitionMatrix(idx,idx+1) = 0.25;
    end
    if i < 4 && idx != 11
      transitionMatrix(idx,idx+4) = 0.25;
    end
  end
end
solution = (eye(14,14) - transitionMatrix)\repmat([-1], 14, 1);
[0 solution(1) solution(2) solution(3);
  solution(4) solution(5) solution(6) solution(7);
  solution(8) solution(9) solution(10) solution(11);
  solution(12) solution(13) solution(14) 0]
