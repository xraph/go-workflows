//! Minimal crate whose Makefile defines `fmt` and `test` but deliberately NOT
//! `clippy`, `audit`, `build` or `docs`. That asymmetry is the point: it proves
//! the workflow's probe resolves each target independently rather than
//! switching wholesale on the Makefile's existence.

/// Returns the sum of `a` and `b`.
pub fn add(a: i64, b: i64) -> i64 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::add;

    #[test]
    fn adds_two_numbers() {
        assert_eq!(add(2, 3), 5);
    }
}
