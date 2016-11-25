transitionMatrix = eye(16) * 0;
idx = 0;
for i = 1:4
  for j = 1:4
    idx++;
    if (i == 1 && j == 1) || (i == 4 && j == 4)
      transitionMatrix(idx,idx) = 1;
      continue
    end
    if j > 1
      transitionMatrix(idx,idx-1) = 0.25;
    else
      transitionMatrix(idx,idx) += 0.25;
    end
    if i > 1
      transitionMatrix(idx,idx-4) = 0.25;
    else
      transitionMatrix(idx,idx) += 0.25;
    end
    if j < 4
      transitionMatrix(idx,idx+1) = 0.25;
    else
      transitionMatrix(idx,idx) += 0.25;
    end
    if i < 4
      transitionMatrix(idx,idx+4) = 0.25;
    else
      transitionMatrix(idx,idx) += 0.25;
    end
  end
end

immediate = repmat([-1], 14, 1);
solution = [0; (eye(14,14) - transitionMatrix(2:15,2:15)) \ immediate; 0];
resMat = eye(4,4)*0;
for i = 1:4
  for j = 1:4
    resMat(i,j) = solution(j+4*(i-1));
  end
end
resMat
