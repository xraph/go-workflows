//! Minimal crate used to smoke-test the reusable Rust CI workflow.
//! It deliberately ships no Makefile, so the workflow exercises its
//! cargo fallback commands rather than make targets.

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
