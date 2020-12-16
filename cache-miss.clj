(def size 2000)

(def arr
  (vec (map vec
            (for [i (range 0 size)]
              (for [j (range 0 size)]
                (rand-int 5))))))

(defn profile []
  (println
   (time
    (let [ret
          (doall
           (for [i (range 0 size)]
             (doall
              (for [j (range 0 size)]
                (+ (nth (nth arr j) i) i j)))))])))
  (println
   (time
    (let [ret
          (doall
           (for [i (range 0 size)]
             (doall
              (for [j (range 0 size)]
                (+ (nth (nth arr i) j) i j)))))]))))
