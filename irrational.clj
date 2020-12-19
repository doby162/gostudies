(defn newt [root-of guess prec]
  (with-precision prec
  (- guess (/ (- (* guess guess) root-of) (* 2 guess)))))

(defn exp [x n]
  (reduce * (repeat n x)))

(defn main [root-of iterations]
  (let [guess (bigdec 1)]
    (spit "sqrt.edn"
          (loop [i 0 g guess]
            (cond
              (> i iterations)
              g
              :else
              (let [prec (exp 2 i)]
                (println (str "calculating (exp 2 "i") ("prec") digits..."))
                (recur (inc i) (newt root-of g prec))))))))

(defn test [] (main 2 25))
