This function is derived from code originally posted by Sat Kumar Tomer
(satkumartomer@gmail.com)
See also: http://vsp.pnnl.gov/help/Vsample/Design_Trend_Mann_Kendall.htm

The purpose of the Mann-Kendall (MK) test (Mann 1945, Kendall 1975, Gilbert
1987) is to statistically assess if there is a monotonic upward or downward
trend of the variable of interest over time. A monotonic upward (downward)
trend means that the variable consistently increases (decreases) through
time, but the trend may or may not be linear. The MK test can be used in
place of a parametric linear regression analysis, which can be used to test
if the slope of the estimated linear regression line is different from
zero. The regression analysis requires that the residuals from the fitted
regression line be normally distributed; an assumption not required by the
MK test, that is, the MK test is a non-parametric (distribution-free) test.
Hirsch, Slack and Smith (1982, page 107) indicate that the MK test is best
viewed as an exploratory analysis and is most appropriately used to
identify stations where changes are significant or of large magnitude and
to quantify these findings.

Input:
    x:   a vector of data
    alpha: significance level (0.05 default)

Output:
    trend: tells the trend (increasing, decreasing or no trend)
    h: True (if trend is present) or False (if trend is absence)
    p: p value of the significance test
    z: normalized test statistics

Example:
```Go
import mk "github.com/detailyang/mann-kendall-go"

rv := mk.Test([]float64{1.1, 1.1, 1.1, 1.1, 1.1}, 0.05)
```